<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-card class="box-card">
        <template #header>
          <div class="card-header">
            <el-input
                placeholder="请输入内容"
                v-model="search"
                class="input-with-select"
            >
              <template #prefix>
                <i class="el-input__icon el-icon-search"></i>
              </template>
            </el-input>
          </div>
        </template>
        <el-scrollbar height="440px" class="infinite-list-wrapper" style="overflow:auto" v-infinite-scroll="load">
          <div v-for="o in 34" :key="o" class="text item">
            {{ '列表内容 ' + o }}
          </div>
        </el-scrollbar>
      </el-card>
    </el-aside>
    <el-main>
      <tool-editor/>
    </el-main>
  </el-container>
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'

import {defineComponent, ref} from 'vue'

export default defineComponent({
  components: {
    ToolEditor
  },

  setup() {
    document.title = "editing"
    return {
      search: ref(''),
    }
  },

  methods: {
    load(): void {
      console.log("load")
    }
  }
})
</script>

<style scoped>
.hidden-xs-only {
  padding-top: 20px;
}
</style>
