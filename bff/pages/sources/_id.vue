<template>
  <div class="root">
    <div class="sidebar">
      <source-list @add="add" :sources="sources"/>
    </div>

    <section class="list">
      <div>
        <ul>
          <!-- TODO key -->
          <article-list-item v-for="item in articles" :key="item.title"
                             :item="item" @read="read"/>
        </ul>
      </div>
    </section>
  </div>
</template>

<script lang="ts">
  import axios from '~/plugins/axios'
  import SourceList from '~/components/SourceList.vue'
  import {Article} from "~/common/Article";
  import ArticleListItem from "~/components/ArticleListItem";

  async function fetchArticles(sourceId: number): Promise<Article[]> {
    let articlesRes = await axios.get(`/api/sources/${sourceId}/articles`);
    return articlesRes.data;
  }

  async function fetchSources(): Promise<string[]> {
    const sourcesRes = await axios.get('/api/sources');
    return sourcesRes.data;
  }

  export default {
    async asyncData({params}) {
      const sourceId = parseInt(params['id']);

      const [articles, sources] =
        await Promise.all([fetchArticles(sourceId), fetchSources()]);

      return {
        sourceId: sourceId,
        articles: articles,
        sources: sources,
      }
    },

    methods: {
      read(item: Article) {
        console.log(item);
        alert(item.title);
      },
      async add(newUrl: string) {
        await axios.post(`/api/sources`, {
          url: newUrl
        });
        this.sources = await fetchSources();
      }
    },

    components: {
      ArticleListItem,
      SourceList,
    }
  }
</script>

<style>
  .root {
    margin: 10px;
    display: flex;
  }
</style>
