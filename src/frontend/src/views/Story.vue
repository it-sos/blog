<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-tabs type="border-card">
        <el-tab-pane label="文章">
          <template #header>
            <div class="card-header">
              <el-input
                  placeholder="请输入内容"
                  v-model="keyword"
                  class="input-with-select"
              >
                <template #prefix>
                  <i class="el-input__icon el-icon-search"></i>
                </template>
              </el-input>
            </div>
          </template>
          <div class="infinite-list-wrapper" style="height:400px;overflow-y:auto;overflow-x: hidden;"
               v-infinite-scroll="load">
            <div v-bind:key="idx" v-for="(art, idx) in article" class="text item dirs">
              <el-link v-right-click="rightMenu" href="javascript:void(0);"><span
                  v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
              <el-tag style="margin-left:0.3rem;" effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
            </div>
          </div>
        </el-tab-pane>
        <el-tab-pane label="专题">
        </el-tab-pane>
        <el-tab-pane label="标签">角色管理</el-tab-pane>
      </el-tabs>
    </el-aside>
    <el-main>
      <tool-editor/>
    </el-main>
  </el-container>
  <right-menu :position="position" ref="childRef" />
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'
import RightMenu from '../components/RightMenu.vue'

import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {Position} from "@/components/RightMenu.vue";
// import {ElMessage} from "element-plus";

export default defineComponent({
  components: {
    ToolEditor,
    RightMenu,
  },

  directives: {
    rightClick: {
      mounted(el, binding) {
        el.addEventListener("contextmenu", function (ev: any) {
          ev.preventDefault();
          binding.value()
          return false;
        }, false)
      }
    }
  },

  setup() {
    document.title = "editing"

    const childRef = ref(null)
    const rightMenu = () => {
      childRef.value.showMenu()
    }

    const article: any[] = []
    const state = reactive({
      keyword: "",
      page: 0,
      size: 10,
      noMore: false,
      article: article,
      show: false,
    });

    const position :Position = reactive({
      x: 0,
      y: 0,
    });

    const defaults = () => {
      state.page = 0
      state.noMore = false
    }

    // const handleCommand = (command: any) => {
    //   ElMessage(`click on item ${command}`);
    // };

    return {
      ...toRefs(state),
      position,
      childRef,
      rightMenu,
      defaults,
    }
  },

  created() {
    let timer: any = null
    watch(
        () => this.keyword,
        () => {
          clearTimeout(timer)
          timer = setTimeout(() => {
            this.defaults()
            this.load()
          }, 1000)
        }
    )
  },

  methods: {
    load(): void {
      if (this.noMore) {
        return
      }
      this.article = [
        {
          article: {
            title: "标准标题1"
          },
          duration: "2秒"
        },
        {
          article: {
            title: "标准标题2"
          },
          duration: "2秒"
        },
      ]
      // this.page++
      // this.$http.get('/article/list', {
      //   params: {
      //     "page": this.page,
      //     "size": this.size,
      //     // "keyword": this.keyword
      //   }
      // }).then((response) => {
      //   if (response.data.length === 0) {
      //     this.noMore = true
      //     return
      //   }
      //   response.data.forEach((v: any) => {
      //     this.article.push(v)
      //   })
      // }).catch((error) => {
      //   console.log(error)
      // })
    }
  }
})
</script>

<style lang="scss" scoped>
.hidden-xs-only {
  padding-top: 20px;
}

.dirs {
  margin-bottom: 8px;
  display: flex;
  justify-content: space-between;

  span {
    font-size: 12px;
  }

}
</style>
