export default function onButtonClick(
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
