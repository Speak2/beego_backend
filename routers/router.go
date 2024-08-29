package routers

import (
	"cats_backend/controllers"
	beego "github.com/beego/beego/v2/server/web"  
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/random-cat", &controllers.RandomCatController{}) 
	beego.Router("/api/breeds", &controllers.BreedsController{}, "get:GetAllBreeds")
	beego.Router("/api/breed-images", &controllers.BreedsController{}, "get:GetBreedImages")
	beego.Router("/api/favorites", &controllers.VotingController{}, "post:AddFavorite")
	beego.Router("/api/votes", &controllers.VotingController{}, "post:Vote")
	beego.Router("/api/get_favorites", &controllers.FavoritesController{}, "get:GetFavorites")
}
