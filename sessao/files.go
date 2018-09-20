package sessao

import (
	"encoding/gob"
	"net/http"
	"os"
)

type handlerFiles struct {
}

func newHandlerFiles() (handler, error) {
	err := os.Mkdir("./sessaoarquivos", 0644)
	if err != nil && !os.IsExist(err) {
		return nil, err
	}
	h := handlerFiles{}
	return h, nil
}

func (h handlerFiles) inicia(r *http.Request) (map[string]interface{}, error) {
	novaSess := make(map[string]interface{})
	cookie, err := r.Cookie(nome)
	if err != nil {
		return novaSess, nil
	}
	f, err := os.Open("./sessaoarquivos/" + cookie.Value + ".txt")
	if err != nil {
		return novaSess, nil
	}
	defer f.Close()
	dec := gob.NewDecoder(f)
	err = dec.Decode(&novaSess)
	return novaSess, err
}

func (h handlerFiles) salva(w http.ResponseWriter, r *http.Request, sess map[string]interface{}) error {
	cookie, err := r.Cookie(nome)
	if err != nil {
		cookie = &http.Cookie{
			Name:  nome,
			Value: stringAleatoria(),
		}
		http.SetCookie(w, cookie)
	}
	f, err := os.OpenFile("./sessaoarquivos/"+cookie.Value+".txt", os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	enc := gob.NewEncoder(f)
	return enc.Encode(sess)
}

func (h handlerFiles) destroi(r *http.Request) error {
	cookie, err := r.Cookie(nome)
	if err != nil {
		return nil
	}
	return os.Remove("./sessaoarquivos/" + cookie.Value + ".txt")
}
