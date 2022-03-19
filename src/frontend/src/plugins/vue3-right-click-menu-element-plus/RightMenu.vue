<template>
  <transition name="slide-fade">
    <div v-if="state.show" class="contextMenu" @mouseleave="hideMenu" @mouseenter="showMenu" :style="positionCss">
      <el-dropdown @command="handleCommand">
        <el-dropdown-menu>
          <el-dropdown-item :key="index" v-for="(m, index) in menu[state.active]" :command="m.command" :icon="m.icon">{{m.title}}</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </transition>
</template>

<script lang="ts">

import {defineComponent, inject, reactive, ref, toRaw} from 'vue'
import {Delete} from "@element-plus/icons-vue"

export default defineComponent({
  name: "RightMenu",

  setup(props, context) {
    let state = reactive({
      position: ref({ x:0, y:0 }),
      show: ref(true),
      active: ref<string>('test'),
      menu: ref({'test':[
          {
            'icon': Delete,
            'title': '需要配置 menu 项，以填充更多功能',
            'command': 'demo',
          },
        ]})
    })
    state = inject('right-click-menu',  state)
    const menu = toRaw(state.menu)

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
      Delete,
      menu,
      state,
      showMenu,
      hideMenu,
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
  z-index: 9999;
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
