/* eslint-disable no-param-reassign */
import { createSlice } from "@reduxjs/toolkit";

const formSlice = createSlice({
  name: "form",
  initialState: {
    name: "",
    birthday: "",
    gender: "",
    genderPreference: "",
    looking: "",
    pronouns: "",
    height: { foot: 0, inch: 0 },
    location: ""
  },
  reducers: {
    setName: (state, action) => {
      state.name = action.payload;
    },
    setBirthday: (state, action) => {
      state.birthday = action.payload;
    },
    setGenderPreference: (state, action) => {
      state.genderPreference = action.payload;
    },
    setGender: (state, action) => {
      state.gender = action.payload;
    },
    setLooking: (state, action) => {
      state.looking = action.payload;
    },
    setPronouns: (state, action) => {
      state.pronouns = action.payload;
    },
    setHeight: (state, action) => {
      state.height = action.payload;
    },
    setLocation: (state, action) => {
      state.location = action.payload;
    }
  }
});

export const {
  setName,
  setBirthday,
  setGenderPreference,
  setGender,
  setLooking,
  setPronouns,
  setHeight,
  setLocation
} = formSlice.actions;

export default formSlice.reducer;
