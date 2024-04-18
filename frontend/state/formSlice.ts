/* eslint-disable no-param-reassign */
import { createSlice } from "@reduxjs/toolkit";

interface Photo {
  filePath: string;
  caption: string;
}

interface Height {
  foot: number;
  inch: number;
}

const formSlice = createSlice({
  name: "form",
  initialState: {
    id: "" as string,
    fullName: "" as string,
    email: "" as string,
    name: "" as string,
    birthday: "" as string,
    gender: "" as string,
    genderPreference: "" as string,
    looking: "" as string,
    pronouns: "" as string,
    height: { foot: 0, inch: 0 } as Height,
    location: "" as string,
    school: "" as string,
    job: "" as string,
    religion: "" as string,
    politics: "" as string,
    drinkHabit: "" as string,
    smokeHabit: "" as string,
    weedHabit: "" as string,
    drugHabit: "" as string,
    passion: [] as string[],
    promptBio: "" as string,
    responseBio: "" as string,
    photos: [] as Photo[],
    instagram: "" as string,
    notifications: false as boolean
  },
  reducers: {
    setFullName: (state, action) => {
      console.log(state.fullName);
      state.fullName = action.payload;
    },
    setEmail: (state, action) => {
      console.log(state.email);
      state.email = action.payload;
    },
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
    },
    setNotifications: (state, action) => {
      state.notifications = action.payload;
    },
    setId: (state, action) => {
      state.id = action.payload;
    }
  }
});

export const {
  setFullName,
  setEmail,
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
  setInstagram,
  setNotifications,
  setId
} = formSlice.actions;

export default formSlice.reducer;
