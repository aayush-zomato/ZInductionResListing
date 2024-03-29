package httphandler

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"Zinduction/reslisting/driver"
	models "Zinduction/reslisting/models"
	repository "Zinduction/reslisting/repository"
	post "Zinduction/reslisting/repository/post"
)

func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL), //creating new Post from DB
	}
}

type Post struct{
	repo repository.PostRepo // repository interface
}

func(p *Post) Fetch(w http.ResponseWriter, r *http.Request){
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondwithJSON(w, http.StatusOK, payload)
}


func(p *Post) Create(w http.ResponseWriter, r *http.Request){
	post := models.Post{}
	json.NewDecoder(r.Body).Decode(&post)

	newID, err := p.repo.Create(r.Context(), &post)
	fmt.Println(newID)
	if err != nil{
		// respondWithError(w, http.StatusInternalServerError, "Server Error") //creating a change
		respondWithError(w, http.StatusInternalServerError, err.Error()) //creating a change
	}
	respondwithJSON(w, http.StatusCreated, map[string]string{"message":"Successfully Created"})
}

func (p *Post) Update(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data := models.Post{ID: int(id)}
	json.NewDecoder(r.Body).Decode(&data)
	payload, err := p.repo.Update(r.Context(), &data)

	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Post) GetByID(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.repo.GetByID(r.Context(), int64(id))

	if err!=nil{
		respondWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

func (p *Post) Delete(w http.ResponseWriter, r *http.Request){
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	_, err := p.repo.Delete(r.Context(), int64(id))

	if err!=nil{
		respondWithError(w, http.StatusInternalServerError, err.Error())
	}
	respondwithJSON(w, http.StatusMovedPermanently, map[string] string{"message":"Delete Successfully"})
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondwithJSON(w, code, map[string]string{"message":msg})
}
