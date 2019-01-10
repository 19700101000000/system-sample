<template lang="pug">
  b-container

    b-card.my-2(
      v-if="$store.state.name"
      title="Upload image"
      no-body)
      div(v-if="imagePath != ''" style="max-height: 300px; overflow: hidden;")
        b-card-img(:src="imagePath" top fluid)
      b-card-body
        h4.card-title Upload image
        p.card-text.text-danger(v-if="errorMsg != ''") {{ errorMsg }}
        b-form(@submit="onSubmit" @reset="onReset")
          b-form-group(
            horizontal
            label="Image:"
            label-class="text-sm-right"
            :label-cols="2")
            b-form-file.pt-2(v-model="imageFile" :state="imagePath !== ''" placeholder="Choose an image..." accept=".jpg, jpeg, .png" @change="selectImage")
          b-form-group(
            horizontal
            label="Categories:"
            label-class="text-sm-right"
            :label-cols="2")
            b-form-checkbox-group.pt-2(v-model="selectedCategories" name="categories" :options="optCategories")
          div.mt-2.text-center
            b-button(type="submit" variant="outline-success" :disabled="imagePath == ''") Send
            b-button.ml-2(type="reset" variant="outline-danger" :disabled="imagePath == ''") Reset
    b-row
      b-col
        rank-table
      b-col
        rank-table
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"
import AdCarousel from "~/components/AdCarousel.vue"
import RankTable from "~/components/RankTable.vue"

@Component({
  components: {
    AdCarousel,
    RankTable
  }
})
export default class extends Vue {
  public imageFile = "";
  public imagePath = "";
  public errorMsg = "";

  public selectedCategories: string[] = [];
  public optCategories = {};

  public onSubmit(e: Event) {
    e.preventDefault();
    if (this.imageFile == "") {
      this.errorMsg = "Please selected image."
      return;
    }
    const url = "/api/upload/image";
    let params = new FormData();
    params.append("image", this.imageFile);
    for (var category of this.selectedCategories) {
      params.append("categories", category);
    }
    axios.post(url, params).then((result) => {

    }).catch(() => {
      this.errorMsg = "Failed connection.";
    });
  }
  public onReset(e: Event) {
    this.imageFile = "";
    this.imagePath = "";
    this.errorMsg = "";
  }
  public selectImage(e: Event) {
    if (e.target instanceof HTMLInputElement && e.target.files && e.target.files[0]) {
      const type = e.target.files[0].type;
      if (type !== "image/jpeg" && type !== "image/png") {
        this.errorMsg = "Please selected image(jpeg or png)";
        this.imageFile = "";
        this.imagePath = "";
        return;
      }
      let reader = new FileReader();
      reader.onload = (e: FileReaderProgressEvent) => {
        if (e.target) {
          this.imagePath = e.target.result;
          this.errorMsg = "";
          window.console.log(e);
        }
      }
      reader.readAsDataURL(e.target.files[0]);
    } else {
      this.imagePath = "";
    }
  }
  public mounted() {
    this.getOptions();
  }
  private getOptions() {
    axios.get("/api/get/categories")
    .then((result) => {
      if (result.data) {
        this.optCategories = result.data.options;
      }
    });
  }
}
</script>

<style lang="sass" scoped></style>
