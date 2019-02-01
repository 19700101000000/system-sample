<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      b-card.mt-2(no-body)
        b-card-body
          h4 MY REQUESTS
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import UserNav from "~/components/UserNav.vue"

@Component({
  components: {
    UserNav,
  }
})
export default class extends Vue {
}
</script>

<style lang="sass"></style>
