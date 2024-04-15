export function onButtonClick(
  value: string,
  buttonValue: string,
  setSelectedButton: Function,
  onChange: Function
) {
  if (value === buttonValue) {
    onChange("");
    setSelectedButton("");
  } else {
    onChange(buttonValue);
    setSelectedButton(buttonValue);
  }
}

export function onButtonClickArray(
  value: string[],
  buttonValue: string,
  setSelectedButtons: Function,
  onChange: Function
) {
  // Check if the buttonValue is already in the array
  if (value.includes(buttonValue)) {
    // If it is, remove it from the array
    const newArray = value.filter((v) => v !== buttonValue);
    onChange(newArray);
    setSelectedButtons(newArray);
  } else if (value.length < 5) {
    // If it's not in the array and the array length is less than 5, add it to the array
    const newArray = [...value, buttonValue];
    onChange(newArray);
    setSelectedButtons(newArray);
  }
}
