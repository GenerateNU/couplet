const calculateAge = (personAge: Date) => {
  const todayDate = new Date();
  const todayDay = todayDate.getDate();
  const todayMonth = todayDate.getMonth();
  const todayYear = todayDate.getFullYear();
  const personDay = personAge.getDate();
  const personMonth = personAge.getMonth();
  const personYear = personAge.getFullYear();
  let age = todayYear - personYear;

  if (todayMonth < personMonth || (todayMonth === personMonth && todayDay < personDay)) {
    age -= 1;
  }
  return age;
};

export default calculateAge;
