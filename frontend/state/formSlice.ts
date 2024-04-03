import { createSlice } from "@reduxjs/toolkit";

const formSlice = createSlice({
  name: "form",
  initialState: {
    name: "",
    birthday : "",
    genderPreference : "",
  },
  reducers: {
    setName: (state, action) => {
      state.name = action.payload;
    },
    setBirthday : (state, action) => {
      state.birthday = action.payload;
    },
    setGenderPreference : (state, action) => {
      state.genderPreference = action.payload
    }
  }
});

export const { setName, setBirthday, setGenderPreference } = formSlice.actions;

export default formSlice.reducer;
