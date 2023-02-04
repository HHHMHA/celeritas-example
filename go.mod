module myapp

go 1.19

replace github.com/HHHMHA/celeritas => ../celeritas

require github.com/HHHMHA/celeritas v0.0.0-20221228181119-4000bbe11291

require (
	github.com/go-chi/chi/v5 v5.0.8 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
