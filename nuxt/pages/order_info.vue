<template lang="pug">
div
  p 受注詳細
</template>

<script>
export default {
  fetch({store, redirect}) {
    if (!store.state.auth) {
      return redirect('/login')
    }
  }
}
</script>
