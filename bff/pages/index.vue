<template>
  <div class="root">
    <div class="sidebar">
      <source-list @add="add" :sources="sources"/>
    </div>
  </div>
</template>

<script lang="ts">
  import axios from '~/plugins/axios'
  import SourceList from '~/components/SourceList.vue'

  export default {
    async asyncData() {
      const pSourcesRes = axios.get('/api/sources');
      const sourcesRes = await pSourcesRes;

      return {
        sources: sourcesRes.data,
      }
    },

    methods: {
      async add(newUrl: string) {
        alert("TODO");
      }
    },

    components: {
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
