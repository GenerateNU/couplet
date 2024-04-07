/* eslint-disable no-param-reassign */
import { createSlice } from "@reduxjs/toolkit";

interface Photo {
  filePath: string;
  caption: string;
}

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
    location: "",
    school: "",
    job: "",
    religion: "",
    politics: "",
    drinkHabit: "",
    smokeHabit: "",
    weedHabit: "",
    drugHabit: "",
    passion: [],
    promptBio: "",
    responseBio: "",
    photos: [] as Photo[],
    instagram: ""
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
    },
    setSchool: (state, action) => {
      state.school = action.payload;
    },
    setJob: (state, action) => {
      state.job = action.payload;
    },
    setReligion: (state, action) => {
      state.religion = action.payload;
    },
    setPolitics: (state, action) => {
      state.politics = action.payload;
    },
    setDrinkHabit: (state, action) => {
      state.drinkHabit = action.payload;
    },
    setSmokeHabit: (state, action) => {
      state.smokeHabit = action.payload;
    },
    setWeedHabit: (state, action) => {
      state.weedHabit = action.payload;
    },
    setDrugHabit: (state, action) => {
      state.drugHabit = action.payload;
    },
    setPassion: (state, action) => {
      state.passion = action.payload;
    },
    setPromptBio: (state, action) => {
      state.promptBio = action.payload;
    },
    setResponseBio: (state, action) => {
      state.responseBio = action.payload;
    },
    setPhotos: (state, action) => {
      state.photos = action.payload;
    },
    setInstagram: (state, action) => {
      state.instagram = action.payload;
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
  setLocation,
  setSchool,
  setJob,
  setReligion,
  setPolitics,
  setDrinkHabit,
  setDrugHabit,
  setSmokeHabit,
  setWeedHabit,
  setPassion,
  setPromptBio,
  setResponseBio,
  setPhotos,
  setInstagram
} = formSlice.actions;

export default formSlice.reducer;
