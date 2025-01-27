package rank

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikkerella/hitotose/gin/model/rank"
	"github.com/nikkerella/hitotose/gin/mw"
	rank_svc "github.com/nikkerella/hitotose/gin/svc/rank"
)

var (
	svc rank_svc.Service
)

func Init(e *gin.Engine) {
	svc = rank_svc.NewService()

	anon := e.Group("/api/rank")
	{
		anon.GET("/", query)
		anon.GET("/query", query)
		anon.GET("/members", members)
	}

	auth := e.Group("/api/rank").Use(mw.Auth)
	{
		auth.POST("/add", add)
		auth.POST("/delete", delete)
		auth.POST("/create", create)
		auth.POST("/remove", remove)
		auth.POST("/sort", sort)
	}
}

func add(c *gin.Context) {
	var reqBody struct {
		RankID string `json:"rank_id"`
		Name   string `json:"name"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	svc.Add(reqBody.RankID, reqBody.Name)
	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func create(c *gin.Context) {
	title := c.PostForm("title")
	svc.Create(title)
	c.Redirect(http.StatusSeeOther, "/rank")
}

func delete(c *gin.Context) {
	id := c.PostForm("id")
	svc.Delete(id)
	c.Redirect(http.StatusSeeOther, "/rank")
}

func members(c *gin.Context) {
	rID := c.Query("rank_id")
	r := svc.ByID(rID)
	c.JSON(http.StatusOK, gin.H{
		"all_members":      r.Members,
		"unranked_members": r.UnrankedMembers(),
		"ranked_members":   r.RankedMembers(),
	})
}

func query(c *gin.Context) {
	// id := c.Query("id")

	// if id != "" {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"rank": svc.ByID(id),
	// 	})
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"ranks": svc.Query(),
	// 	})
	// }

	c.JSON(http.StatusOK, gin.H{
		"ranks": svc.Query(),
	})
}

func remove(c *gin.Context) {
	var reqBody struct {
		RankID string `json:"rank_id"`
		Name   string `json:"name"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rank := svc.ByID(reqBody.RankID)
	rank.Remove(reqBody.Name)
	svc.Update(rank)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func sort(c *gin.Context) {
	var reqBody struct {
		RankID          string        `json:"rank_id"`
		UnrankedMembers []rank.Member `json:"unranked_members"`
		RankedMembers   []rank.Member `json:"ranked_members"`
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var ranked []rank.Member
	for i, member := range reqBody.RankedMembers {
		// r.Update(i+1, member.Name)
		ranked = append(ranked, rank.Member{
			Order: i + 1,
			Name:  member.Name,
		})
	}

	var unranked []rank.Member
	for _, member := range reqBody.UnrankedMembers {
		unranked = append(unranked, rank.Member{
			Order: 0,
			Name:  member.Name,
		})
	}

	r := svc.ByID(reqBody.RankID)
	r.Members = ranked
	r.Members = append(r.Members, unranked...)

	svc.Update(r)

	c.JSON(http.StatusOK, gin.H{
		"msg": "OK",
	})
}

func update(c *gin.Context) {
	id := c.PostForm("id")
	newTitle := c.PostForm("title")

	rank := svc.ByID(id)
	rank.Title = newTitle
	svc.Update(rank)
}
