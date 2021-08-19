<template>
  <transition name="slide-fade">
    <div v-if="state.show" class="contextMenu" @mouseleave="hide" :style="positionCss">
      <el-dropdown @command="handleCommand">
        <el-dropdown-menu>
          <el-dropdown-item>黄金糕</el-dropdown-item>
          <el-dropdown-item>狮子头</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </div>
  </transition>
</template>

<script lang="ts">

import {defineComponent, PropType, reactive} from 'vue'

export interface Position {
  x: number
  y: number
}

export default defineComponent({
  name: "RightMenu",
  props: {
    position: {
      type: Object as PropType<Position>,
      required: true,
    },
  },

  setup() {
    const state = reactive({
      show: true
    })

    const showMenu = () => {
      state.show = true
    }

    const hideMenu = () => {
      state.show = false
    }

    return {
      state,
      showMenu,
      hideMenu,
    }
  },

  computed: {
    positionCss: function(): string {
      return `left: ${this.position.x}px; top: ${this.position.y}px`
    },
  },

  methods: {
    handleCommand(command: any) {
      console.log('click on item ' + command);
    },

    hide() {
      this.state.show = false
      // this.showCss = false
      console.log("hidden.")
    }
  }
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