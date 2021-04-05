<template>
  <form @submit.prevent="onSubmit">
    <p>
      <label for="name">Name:</label>
      <input id="name" v-model="person.name" placeholder="name" />
    </p>
    <p>
      <input
        type="radio"
        id="male"
        v-model="person.gender"
        name="gender"
        value="male"
      />
      <label for="male">Male</label><br />
      <input
        type="radio"
        id="female"
        v-model="person.gender"
        name="gender"
        value="female"
      />
      <label for="female">Female</label><br />
      <input
        type="radio"
        id="other"
        v-model="person.gender"
        name="gender"
        value="other"
      />
      <label for="other">Other</label>
    </p>
    <p>
      <input type="submit" value="Submit" />
    </p>
  </form>
</template>

<script>
export default {
  data() {
    return {
      person: {
        name: null,
        gender: null,
      },
    };
  },
  methods: {
    onSubmit(e) {
      console.log("event", e);
      console.log("person model", this.person);
    },
  },
};
</script>

<style>
</style>
