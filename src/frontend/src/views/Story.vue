<template>
  <el-drawer v-model="drawer" :with-header="false">
    <el-card class="box-card">
      <template #header>
        <div class="card-header">
          <el-input
              placeholder="请输入内容"
              v-model.trim="keyword"
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
        <div v-bind:key="art.id" v-for="art in article" class="text item dirs">
          <el-link href="javascript:void(0);" @click="edit(art.id)"
                   v-right-click="rightMenuFunc('article', art.id)"><span
              v-html="art.title_match ? art.title_match : art.title"></span></el-link>
          <span style="margin: 0 5px;color: #303133" v-if="art.is_state === 1"><el-icon><Lock/></el-icon></span>
          <el-tag style="margin-left:0.3rem;" effect="plain" type="warning" size="mini">{{ art.duration }}</el-tag>
        </div>
      </div>
    </el-card>
  </el-drawer>
  <el-container>
    <el-aside style="width:30px;padding-top:20px;overflow:hidden;">
      <el-icon style="font-size: 32px;color:#555;cursor:pointer;" @click="drawer = true">
        <Fold v-if="!drawer"/>
        <Expand v-else/>
      </el-icon>
    </el-aside>
    <el-main>
      <tool-editor @syncArticleList="syncArticleList"/>
    </el-main>
  </el-container>
  <right-menu @trigger="rightMenuTrigger"/>
</template>
<script lang="ts">

import {computed, defineComponent, inject, provide, reactive, ref, toRefs, watch} from 'vue'
import {ElMessage} from "element-plus";
import ToolEditor from "../components/ToolEditor.vue";
import {CirclePlus, Delete, Expand, Fold, Lock} from "@element-plus/icons-vue";
import {onBeforeRouteLeave, onBeforeRouteUpdate, useRouter} from 'vue-router';
import {useStore} from "../store/store";

export default defineComponent({
  components: {
    ToolEditor,
    Fold,
    Expand,
    Lock,
  },

  setup() {
    document.title = "editing"
    const $axios: any = inject('$axios')
    const router = useRouter()
    const store = useStore()
    store.commit('setIsBackend', true)

    const confirmLevel = () => {
      if (!store.getters.getSaved()) {
        const answer = window.confirm(
            '你确认要离开吗？将丢失未保存的内容！'
        )
        // 取消导航并停留在同一页面上
        if (!answer) return false
      }
    }

    onBeforeRouteUpdate((to, from) => confirmLevel())
    onBeforeRouteLeave((to, from) => confirmLevel())

    const state = reactive({
      keyword: ref(""),
      page: ref(0),
      size: ref(50),
      noMore: ref(false),
    });

    let timer: any = null
    watch(() => state.keyword, () => {
      clearTimeout(timer)
      timer = setTimeout(articleListRest, 1000)
    })

    const load = () => {
      if (state.noMore) {
        return
      }
      state.page++
      $axios.get('/admin/articles', {
        responseType: "json",
        params: {page: state.page, size: state.size, keyword: state.keyword},
      }).then((response: any) => {
        response.data.forEach((v: ArticleList) => {
          store.commit('article/append', v)
        })
        if (response.data.length < state.size) {
          state.noMore = true
        }
      }).catch((error: any) => {
        ElMessage.warning(error.response.data.message)
      })
    }

    // 重置文章列表
    const articleListRest = () => {
      state.page = 0
      state.noMore = false
      store.commit('article/reset')
      load()
    }

    // eslint-disable-next-line
    const enum OPT_TYPE {
      // eslint-disable-next-line
      Add = "add",
      // eslint-disable-next-line
      Update = "update"
    }

    // 编辑文章并动态同步文章列表
    const syncArticleList = (type: OPT_TYPE, id: number, title: string, is_state: number) => {
      // 更新
      if (type == OPT_TYPE.Update) {
        store.commit('article/update', {id: id, title: title, is_state: is_state})
      }
      // 新增
      else if (type == OPT_TYPE.Add) {
        store.commit('article/add', {id: id, title: title, is_state: is_state})
      }
    }

    // 删除操作
    const deleteConfirm = (id: number) => {
      store.commit('article/remove', {id: id})
    };

    let rightMenu = reactive({
      id: ref(0),
      position: ref({x: 0, y: 0}),
      show: ref(false),
      active: ref('article'),
      menu: ref({
        'article': [
          {
            'icon': CirclePlus,
            'title': '新建',
            'command': 'article_add',
          },
          // {
          //   'icon': 'el-icon-edit',
          //   'title': '编辑',
          //   'command': 'article_edit',
          // },
          {
            'icon': Delete,
            'title': '删除',
            'command': 'article_delete',
          },
        ],
      })
    })
    provide('right-click-menu', rightMenu)

    const rightMenuFunc = (active: string, id: number) => {
      return () => {
        rightMenu.active = active
        rightMenu.id = id
        return rightMenu
      }
    }

    const rightMenuTrigger = (type: string) => {
      switch (type) {
        case 'article_add':
          router.push('/e/')
          break;
          // case 'article_edit':
          //   router.push('/e/'+rightMenu.id)
          //   break;
        case 'article_delete':
          deleteConfirm(rightMenu.id)
          break;
      }
    }
    const edit = (id: number) => {
      router.push('/e/' + id)
    }

    return {
      ...toRefs(state),
      drawer: ref(false),
      edit,
      load,
      syncArticleList,
      rightMenu,
      rightMenuFunc,
      rightMenuTrigger,
      article: computed(() => store.getters["article/list"]),
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
