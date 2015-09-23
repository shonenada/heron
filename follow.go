package heron

import (
	"net/http"

	"github.com/shonenada/heron/models"
	"github.com/zenazn/goji/web"
)

func FollowsController(c web.C, w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	db := GetDatabase()
	currentUser, err := CurrentUser(r)
	if err != nil || currentUser.Username == "" {
		RenderJson(w, map[string]interface{}{
			"success": false,
			"message": "user not login",
		})
		return
	}
	if r.Method == "POST" {
		userId := r.PostForm.Get("user_id")
		account := models.Account{}
		db.Where("id = ?", userId).First(&account)
		if account.Username == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "user not found",
			})
			return
		}
		follow := models.Follow{UserId: currentUser.ID, FollowId: account.ID}
		if db.NewRecord(&follow) {
			db.Create(&follow)
			RenderJson(w, map[string]interface{}{
				"success": true,
			})
			return
		} else {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "cannot follow",
			})
		}
		return
	}
}

func FollowController(c web.C, w http.ResponseWriter, r *http.Request) {
	db := GetDatabase()
	currentUser, err := CurrentUser(r)
	if err != nil || currentUser.Username == "" {
		RenderJson(w, map[string]interface{}{
			"success": false,
			"message": "user not login",
		})
		return
	}
	if r.Method == "DELETE" {
		followId := c.URLParams["fid"]
		account := models.Account{}
		db.Where("id = ?", followId).First(&account)
		if account.Username == "" {
			RenderJson(w, map[string]interface{}{
				"success": false,
				"message": "user not found",
			})
			return
		} else {
			follow := models.Follow{}
			db.Where("id = ?", followId).First(&follow)
			if follow.UserId <= 0 {
				RenderJson(w, map[string]interface{}{
					"success": false,
					"message": "not found",
				})
				return
			}
			if follow.UserId != currentUser.ID {
				RenderJson(w, map[string]interface{}{
					"success": false,
					"message": "permission denied",
				})
				return
			}
			db.Delete(&follow)
			RenderJson(w, map[string]interface{}{
				"success": true,
			})
			return
		}
	}
}
