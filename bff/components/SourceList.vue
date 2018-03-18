<template>
  <nav class="source-list">
    ソース一覧
    <ul>
      <li v-for="source in sources">
        <nuxt-link :to="{ name: 'sources-id', params: { id: source.id }}">{{source.title}}</nuxt-link>
      </li>
    </ul>

    <div>
      <input type="text" v-model="newUrl" placeholder="URL">
      <button @click.prevent="add(newUrl)">追加</button>
    </div>
  </nav>
</template>
<script lang="ts">
  import axios from '~/plugins/axios'

  export default {
    props: ['sources'],

    data() {
      return {
        newUrl: "https://www.engadget.com/rss.xml",
      }
    },
    methods: {
      add(newUrl) {
        this.$emit("add", newUrl);
      }
    }
  }
</script>

<style>
  .source-list {
    border: 1px solid black;
    width: 300px;
  }
</style>
