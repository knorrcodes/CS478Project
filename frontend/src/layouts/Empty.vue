<template>
  <div class="home">
    <header>
      <img class="logoImage" src="@/assets/biglogo.png" alt="big team logo" />
    </header>

    <div class="container">
      <div class="row">
        <div class="col-12">
          <slot />
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.logoImage {
  margin-left: 5rem;
  margin-right: 5rem;
}
</style>
