<template>
  <tool-editor/>
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'
import {Options, Vue} from "vue-class-component";

@Options({
  components: {
    ToolEditor
  },
})
export default class Story extends Vue {}
</script>

<style scoped>

</style>