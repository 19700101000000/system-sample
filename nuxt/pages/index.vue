<template lang="pug">
  b-container
    upload-image
    b-row
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import UploadImage from "~/components/UploadImage.vue"

@Component({
  components: {
    UploadImage,
  }
})
export default class extends Vue {
}
</script>

<style lang="sass" scoped></style>
