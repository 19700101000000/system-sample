<template lang="pug">
div
  p foo
</template>

<script>
export default {
  fetch({ store, redirect }) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  }
}
</script>
