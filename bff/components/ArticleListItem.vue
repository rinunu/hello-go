<template>
  <li class="article">
    <div class="article__title">
      <a :href="item.link">{{item.title}}</a>
    </div>
    <!-- TODO セキュアに -->
    <div v-html="item.description">
    </div>
    <!--<button @click.stop="read(item)">読んだ</button>-->
  </li>
</template>

<script lang="ts">
  import {Article} from "~/common/Article";

  export default {
    props: ['item'],
    methods: {
      read: function (item: Article) {
        this.$emit('read', item);
      }
    }
  }

</script>

<style>
  .article {
    margin: 10px;
    border: 1px solid black;
    padding: 10px;
  }

  .article__title {
    margin: 10px 0 20px;
  }

  .article img {
    width: 100%;
  }
</style>
