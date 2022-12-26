package rest

import (
	_admin "backend/rest/admin"
	_auth "backend/rest/auth"
	_user "backend/rest/user"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func RunAPI(address string) error {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding","Authorization" , "Authorization,X-CSRF-Token"},
		AllowCredentials: true,
	}))

	authH := _auth.NewHandler()
	r.Use(authH.SetAuthority)

	auth := r.Group("/auth")
	{
		auth.POST("/signin", authH.SignIn)
		auth.POST("/signout", authH.SignOut)
	}

	admin := r.Group("/admin")
	admin.Use(authH.CheckAdmin)
	{
		h := _admin.NewHandler()

		account := admin.Group("/account")
		{
			account.GET("/list", h.GetAllUserList)

			registration := account.Group("/registration")
			{
				registration.GET("/requests", h.GetAllRegisterRequests)
				registration.POST("/response", h.HandleUserRegistration)
			}
		}

		mileage := admin.Group("/mileage")
		{
			withdrawal := mileage.Group("/withdrawal")
			{
				withdrawal.GET("/requests", h.GetAllMileageRequests)
				withdrawal.POST("/response", h.HandleMileageRequest)
			}

			daily := mileage.Group("/daily")
			{
				daily.POST("/extract", h.PropagateDailyMileageByFile)
				daily.POST("/execute", h.PropagateDailyMileage)
				daily.POST("/initialize", h.InitAllMileage)
			}
			mileage.POST("/update", h.UpdateUserMileage)
		}
	}

	user := r.Group("/user")
	user.Use(authH.CheckUser)
	{
		h := _user.NewHandler()

		account := user.Group(("/account"))
		{
			register := account.Group("/register")
			{
				register.POST("/checkidvalid", h.CheckIdValid)
				register.POST("/request", h.Register)
			}
			account.GET("/resin", h.Resign, authH.SignOut)
			account.GET("", h.GetUser)

			personalinfo := account.Group("/personalinfo")
			{
				personalinfo.GET("", h.GetPersonalInfo)
				personalinfo.POST("", h.ChangePersonalInfo)
			}
		}

		mileage := user.Group("/mileage")
		{
			mileage.GET("/saved", h.GetMileage)
			mileage.GET("/weekly", h.GetWeeklyMileage)
			mileage.POST("/save", h.SaveMileage)

			withdrawal := mileage.Group("/withdrawal")
			{
				withdrawal.POST("/request", h.RequestMileageWithdrawal)
				withdrawal.GET("/status", h.GetMileageWithdrawalStatus)
			}

			friend := mileage.Group("friend")
			{
				friend.GET("/directly", h.GetAllFriendsDailyMileage)
				friend.GET("/bydegree", h.GetAllMileageByDegree)
			}
		}
	}

	return r.Run(address)
}