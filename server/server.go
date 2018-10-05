package server

func Init() {
	r := SetupRouter()
	r.Run(":9527")
}
