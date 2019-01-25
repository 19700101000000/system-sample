<template lang="pug">
  b-container
    b-card.my-1.mx-auto(
      bg-variant="light"
      style="max-width: 30rem"
      no-body)
      b-card-body
        h4 Sign in
        b-form-group(
          horizontal
          label="Name:"
          label-class="text-sm-right"
          label-for="signin-name")
          b-form-input#signin-name(
            v-model="name"
            :state="isName"
            :disabled="isConnecting"
            type="text"
            placeholder="YOUR NAME")

        b-form-group(
          horizontal
          label="Password:"
          label-class="text-sm-right"
          label-for="signin-pass")
          b-form-input#signin-pass(
            v-model="pass"
            :state="isPass"
            :disabled="isConnecting"
            type="password"
            placeholder="YOUR PASSWORD")

        p.text-center.text-danger(v-if="isError") {{ error }}

        p(v-if="isConnecting" style="height: 2.5rem;")
          b-progress.mb-2(:max="1" height="1.5rem")
            b-progress-bar(:value="1" animated striped label="now connecting...")
        p.text-center(v-else style="height: 2.5rem;")
          b-button(
            variant="outline-primary"
            :disabled="(!isName || !isPass)"
            @click="onSubmit") Sign in

        b-button(variant="outline-secondary" block disabled) Sign in from Twitter

        b-button(href="/signup" variant="link" block) Create an acount
</template>

<script lang="ts">
import {
  Component,
  Vue
} from "nuxt-property-decorator"
import axios from "axios"

@Component({})
export default class extends Vue {
  public name: string = "";
  public pass: string = "";
  public error: string = "";
  public isConnecting: boolean = false;
  public isError: boolean = false;

  /* computeds */
  public get isName(): boolean {
    return this.name.length > 0;
  }
  public get isPass(): boolean {
    return this.pass.length > 0;
  }
  /* methods */
  public onSubmit(): void {
    this.toggleConnecting();
    let url = "/api/auth/signin";
    window.setTimeout(() => {
      axios.post(url, {
        name: this.name,
        pass: this.pass
      })
      .then((result) => {
        window.console.log(result);
        let status: string = result.data.status;
        if(status === "ok") {
          window.document.location.href = "/";
        } else {
          this.showErrorMsg("Name or pass is wrong.");
        }
        this.toggleConnecting();
      })
      .catch(() => {
        this.showErrorMsg("Connection failed.");
        this.toggleConnecting();
      });
    }, 1600);
  }
  private toggleConnecting(): void {
    this.isConnecting = !this.isConnecting;
  }
  private showErrorMsg(msg: string): void {
        this.isError = true;
        this.error   = msg;
  }
}
</script>
