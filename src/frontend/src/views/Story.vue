<template>
  <el-container>
    <el-aside width="200px" class="hidden-xs-only">
      <el-tabs type="border-card" class="user-pane">
        <el-tab-pane label="文章">
          <el-card class="box-card">
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
                <el-link v-right-click="rightMenu(1, 'doc')" href="javascript:void(0);"><span
                    v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
                <el-tag style="margin-left:0.3rem;" effect="plain" type="danger" size="mini">{{ art.duration }}</el-tag>
              </div>
            </div>
          </el-card>
        </el-tab-pane>
        <el-tab-pane label="专题">
          <el-card class="box-card">
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
                <el-link v-right-click="rightMenu(1, 'topic')" href="javascript:void(0);"><span
                    v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
              </div>
            </div>
          </el-card>
        </el-tab-pane>
        <el-tab-pane label="标签">
          <el-card class="box-card">
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
                <el-link v-right-click="rightMenu(1, 'tag')" href="javascript:void(0);"><span
                    v-html="art.article.title_match ? art.article.title_match : art.article.title"></span></el-link>
              </div>
            </div>
          </el-card>
        </el-tab-pane>
      </el-tabs>
    </el-aside>
    <el-main>
      <tool-editor/>
    </el-main>
  </el-container>
  <right-menu ref="rightMenuRef" @trigger="trigger" />
</template>
<script lang="ts">

import ToolEditor from '../components/ToolEditor.vue'
import RightMenu from '../components/RightMenu.vue'

import {defineComponent, reactive, ref, toRefs, watch} from 'vue'
import {Position} from "@/components/RightMenu.vue";
import {ElMessage} from "element-plus";
import { ElMessageBox } from 'element-plus';

interface TriggerState {
  id: number
  type: string
}

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
          const position: Position = {
            x: ev.x,
            y: ev.y,
          }
          binding.value(position)
          return false;
        }, false)
      }
    }
  },

  setup() {
    document.title = "editing"

    const rightMenuRef = ref()
    const rightMenu = (id: number, type: string) => {
      return (position: Position) => {
        triggerState = {id, type}
        rightMenuRef.value.showMenu()
        rightMenuRef.value.positionMenu(position)
        document.onclick = () => {
          rightMenuRef.value.hideMenu()
        }
      }
    }

    let article: any[] = []
    let triggerState: TriggerState = {
      id: 0,
      type: ""
    }
    const state = reactive({
      triggerState: triggerState,
      keyword: "",
      page: 0,
      size: 10,
      noMore: false,
      article: article,
    });

    const defaults = () => {
      state.page = 0
      state.noMore = false
    }

    const open = () => {
      ElMessageBox.confirm('此操作将永久删除该文件, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning',
      })
          .then(() => {
            ElMessage({
              type: 'success',
              message: '删除成功!',
            });
          })
          .catch(() => {
            ElMessage({
              type: 'info',
              message: '已取消删除',
            });
          });
    };

    // const handleCommand = (command: any) => {
    //   ElMessage(`click on item ${command}`);
    // };

    const trigger = (type: string) => {
      if (type == "delete") {
        open()
      }
      console.log(type, triggerState)
    }

    return {
      ...toRefs(state),
      rightMenuRef,
      rightMenu,
      defaults,
      trigger,
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

.el-tabs--border-card > .el-tabs__content {
  .el-card {
    border: none;
    box-shadow: none;
  }
}

</style>
<style lang="scss">
.user-pane > .el-tabs__content {
  padding: 0;
}
</style>