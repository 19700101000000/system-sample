<template lang="pug">
b-container
  b-row
    b-col(sm="4")
      user-nav.mt-2
    b-col(sm="8")
      wanted-list-item.mt-2(v-for="wanted in wantedlies" :value="wanted")
</template>

<script lang="ts">
import {
  Component,
  Vue
} from 'nuxt-property-decorator'
import axios from "axios"
import UserNav from "~/components/UserNav.vue"
import WantedListItem from "~/components/WantedListItem.vue"

@Component({
  components: {
    UserNav,
    WantedListItem
  }
})
export default class extends Vue {
  public wantedlies = [];
  public mounted() {
    axios.get("/api/get/works/wantedlies/" + this.$route.params.id).then(({ data }) => {
      this.wantedlies = data.wantedlies;
    });
  }
}
</script>

<style lang="sass"></style>
