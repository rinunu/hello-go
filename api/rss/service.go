package rss

import (
	"github.com/jmoiron/sqlx"
	"github.com/rinunu/hello-go/api/util/reader"
	"database/sql"
	"errors"
)

type Source struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Url   string `json:"url"`
}

type Article struct {
	Id          int
	SourceId    int    `db:"source_id"`
	Title       string `json:"title"`
	Link        string `json:"link"`
	Description string `json:"description"`
}

type RssService struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) *RssService {
	return &RssService{db}
}

// 新規購読を行う
// すでに購読済みなら false を返す
func (service *RssService) Subscribe(url string) (bool, error) {
	source := Source{}
	err := service.db.Get(&source, "SELECT * FROM sources WHERE url = ?", url)
	println(source.Title)
	if err == sql.ErrNoRows {
		// OK
	} else if err != nil {
		println("error")
		println(err.Error())
		return false, err
	} else {
		println("already")
		return false, nil
	}

	rss, err := reader.Read(url)
	if err != nil {
		println("rss read error: " + err.Error())
		return false, err
	}

	service.db.MustExec(`INSERT INTO sources (title, url) VALUES (?, ?)`,
		rss.Title, url)
	return true, nil
}

func (service *RssService) FindAllSources() []Source {
	var sources []Source
	err := service.db.Select(&sources, "SELECT * FROM sources")
	if err != nil {
		panic(err)
	}

	return sources;
}

func (service *RssService) FindArticles(sourceId int) ([]Article, error) {
	var source = service.findSource(sourceId)
	if source == nil {
		return nil, errors.New("no source")
	}
	var articles []Article
	err := service.db.Select(&articles, "SELECT * FROM articles WHERE source_id = ?", sourceId)
	if err != nil {
		panic(err)
	}

	if len(articles) != 0 {
		return articles, nil
	}

	rss, err := reader.Read(source.Url)
	if err != nil {
		return nil, err
	}

	for _, i := range rss.Items {
		newArticle := Article{
			SourceId:    sourceId,
			Title:       i.Title,
			Link:        i.Link,
			Description: i.Description,
		}
		articles = append(articles, newArticle)
		service.db.MustExec(`INSERT INTO articles 
			(source_id, link, title, description)
			VALUES (?, ?, ?, ?)`,
			newArticle.SourceId, newArticle.Link, newArticle.Title, newArticle.Description)
	}
	return articles, nil
}

func (service *RssService) findSource(sourceId int) *Source {
	source := Source{}
	err := service.db.Get(&source, "SELECT * FROM sources WHERE id = ?", sourceId)
	if err == sql.ErrNoRows {
		return nil
	} else if err != nil {
		panic(err)
	}
	return &source;
}
