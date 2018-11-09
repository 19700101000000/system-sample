<template lang="pug">
div
  p 出荷詳細
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
