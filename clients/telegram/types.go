package main

type updateResponse struct {
	Ok     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

func main() {

}
