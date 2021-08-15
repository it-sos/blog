<template>
  <button
      class="menu-item"
      :class="{ 'is-active': isActive ? isActive(): null }"
      @click="action"
      :title="title"
  >
  <i :class="icon"></i>
  </button>
</template>

<script lang="ts">

import { defineComponent } from 'vue'

export default defineComponent({
  name: "MenuItem",
  props: {
    icon: {
      type: String,
      required: true,
    },

    title: {
      type: String,
      required: true,
    },

    action: {
      type: Function,
      required: true,
    },

    isActive: {
      type: Function,
      default: null,
    },
  },
})
</script>

<style lang="scss" scoped>
.menu-item {
  width: 1.75rem;
  height: 1.75rem;
  color: #0D0D0D;
  background-color: transparent;
  border-radius: 0.4rem;
  padding: 0.25rem;
  margin-right: 0.25rem;
  border:1px solid #eaeaea;

&.is-active,
&:hover {
   color: #FFF;
   background-color: #0D0D0D;
 }
}
</style>
