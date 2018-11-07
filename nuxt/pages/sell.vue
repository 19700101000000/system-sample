<template lang="pug">
div
  p bar
</template>

<script>
export default {
  fetch({ store, redirect }) {
    if (!store.state.auth) {
      return redirect('/')
    }
  }
}
</script>

