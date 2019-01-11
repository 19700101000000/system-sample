<template lang="pug">
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
          b-form-file.pt-2(v-if="viewInputImage" v-model="imageFile" :state="imagePath !== ''" placeholder="Choose an image..." accept=".jpg, jpeg, .png" @change="selectImage" :disabled="submitting")
        b-form-group(
          horizontal
          label="Categories:"
          label-class="text-sm-right"
          :label-cols="2")
          b-form-checkbox-group.pt-2(v-model="selectedCategories" name="categories" :options="optCategories" :disabled="submitting")

        div.mt-2(v-if="submitting")
          b-progress.mb-2(:max="1" height="1.5rem")
            b-progress-bar(:value="1" animated striped label="now sending...")
        div.mt-2.text-center(v-else)
          b-button(type="submit" variant="outline-success" :disabled="imagePath == '' || submitting") Send
          b-button.ml-2(type="reset" variant="outline-danger" :disabled="imagePath == '' || submitting") Reset
</template>
<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"

@Component({})
export default class FormImage extends Vue {
  public imageFile = "";
  public imagePath = "";
  public errorMsg = "";

  public viewInputImage = true;
  public submitting = false;

  public selectedCategories: string[] = [];
  public optCategories = {};

  public onSubmit(e: Event) {
    e.preventDefault();
    this.submitting = true;
    window.setTimeout(() => {
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
        if (result.data.status === "ok") {
          this.reset();
        }
        this.submitting = false;
      }).catch(() => {
        this.errorMsg = "Failed connection.";
        this.submitting = false;
      });
    }, 1000);
  }
  public onReset(e: Event) {
    this.reset();
  }
  private reset() {
    this.imageFile = "";
    this.imagePath = "";
    this.errorMsg = "";
    this.selectedCategories = [];
    this.viewInputImage = false;
    this.$nextTick(() => {
      this.viewInputImage = true;
    });
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
