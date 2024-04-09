import COLORS from "../../colors";

const onboardingStyles = {
  TopUiContainer: {
    alignItems: "center",
    flex: 0.35
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium",
    color : COLORS.darkGray
  },
  container: {
    flex: 1,
    marginTop: 34,
    marginBottom: 36
  },
  button: {
    marginBottom: 16
  },
  helperContainer: {
    marginTop: 16
  },
  avoidContainer: {
    flex: 1
  },
  textContainer: {
    padding: 8
  },
  textInputWrapper: {
    marginBottom: 8
  },
  textInput: {
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "#9EA3A2",
    color: "#000000",
    borderRadius: 5,
    padding: 8,
    fontFamily: "DMSansRegular"
  },
  inputWrapper: {
    marginTop: 16
  }
};

export default onboardingStyles;
