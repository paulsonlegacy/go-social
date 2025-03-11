# Notes

- The default mux from the net/http library i.e http.NewServeMux() does not support dynamic URLs and wildcards hence the need for some external mux/router library.

- Gorilla mux library is depreciated and go-chi is used for this project as a modern choice.