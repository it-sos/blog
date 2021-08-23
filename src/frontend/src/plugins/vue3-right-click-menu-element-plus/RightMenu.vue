<template>
  <transition name="slide-fade">
    <div v-if="state.show" class="contextMenu" @mouseleave="hideMenu" @mouseenter="showMenu" :style="positionCss">
      <el-dropdown @command="handleCommand">
        <el-dropdown-menu>
          <el-dropdown-item command="add" icon="el-icon-document-add">新建</el-dropdown-item>
          <el-dropdown-item command="edit" icon="el-icon-edit">编辑</el-dropdown-item>
          <el-dropdown-item command="delete" icon="el-icon-delete">删除</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </transition>
</template>

<script lang="ts">

import {defineComponent, reactive} from 'vue'
import {Position} from "@/plugins/vue3-right-click-menu-element-plus/RightMenu";

export default defineComponent({
  name: "RightMenu",

  setup(props, context) {
    const state = reactive({
      show: false,
      position : {x: 0, y: 0},
    })

    const positionMenu = (position: Position) => {
      state.position = position
    }

    let timer: any = null;
    const hideMenu = () => {
      timer = setTimeout(()=>{
        state.show = false
      }, 500)
    }

    const showMenu = () => {
      clearTimeout(timer)
      state.show = true
    }

    const handleCommand = (command: any) => {
      context.emit('trigger', command)
    }

    return {
      state,
      showMenu,
      hideMenu,
      positionMenu,
      handleCommand,
    }
  },

  computed: {
    positionCss: function(): string {
      return `left: ${this.state.position.x}px; top: ${this.state.position.y}px`
    },
  },
})
</script>

<style scoped>
.contextMenu {
  position: fixed;
  z-index: 1000;
  border: #E4E7ED 1px solid;
  border-radius: 4px;
}

.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.8s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>
