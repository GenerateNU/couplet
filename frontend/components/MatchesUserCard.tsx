import React, { useState } from "react";
import { Text, View, Image, TouchableOpacity, Modal, StyleSheet } from "react-native";
import { MatchesUser } from "./MatchesUserSection";

type MatchesUserCardProps = {
  profile: MatchesUser;
}

const AGE = require("../assets/birthday.png");
const LOCATION = require("../assets/location.png");

const MatchesUserCard: React.FC<MatchesUserCardProps> = ({ profile }) => {
  const [modalVisible, setModalVisible] = useState(false);
  
  const toggleModal = () => {
    setModalVisible(!modalVisible);
  };

  return (
    <TouchableOpacity onPress={toggleModal}>
      <View style={styles.cardContainer}>
        <View style={styles.cardImage}/>
        <View>
          <Text style={styles.cardName}> {profile.name} </Text>
        </View>
      </View>
      <Modal
        animationType="slide"
        transparent={true}
        visible={modalVisible}
        onRequestClose={toggleModal}
      >
        <View style={styles.modalContainer}>
          <TouchableOpacity onPress={toggleModal} style={styles.closeButton}>
            <Text style={styles.closeButtonText}>X</Text>
          </TouchableOpacity>
          <View style={styles.modalContent}>
            <Text style={styles.modalTitle}>{profile.name}</Text>
            <View style={styles.modalValues}>
              <View style={styles.valueContainer}>
                <Image source={AGE} style={styles.icon} />
                <Text style={styles.valueText}>{profile.birthday}</Text>
              </View>
              <View style={styles.valueContainer}>
                <Image source={LOCATION} style={styles.icon} />
                <Text style={styles.valueText}>{profile.location}</Text>
              </View>
            </View>
          </View>
          <View style={styles.profilePicture} />
        </View>
      </Modal>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  cardContainer: {
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 10,
    marginRight: 10
  },
  cardImage: {
    width: 130,
    height: 110,
    backgroundColor: "rgb(200,200,200)",
    borderTopLeftRadius: 10,
    borderTopRightRadius: 10
  },
  cardName: { 
    padding: 10, 
    fontSize: 18 
  },
  modalContainer: {
    flex: 1,
    backgroundColor: "white",
    alignItems: "center",
    paddingHorizontal: 10,
    paddingVertical: 50,
  },
  modalContent: {
    alignItems: "center",
  },
  modalTitle: {
    fontSize: 36,
    fontWeight: "bold",
  },
  modalValues: {
    flexDirection: "row",
    alignItems: "center",
    marginVertical: 10,
  },
  valueContainer: {
    flexDirection: 'row',
    alignItems: 'center',
    marginHorizontal: 25,
  },
  icon: {
    width: 20,
    height: 20,
    resizeMode: 'contain',
    marginRight: 5,
  },
  valueText: {
    fontSize: 16,
  },
  modalName: {
    fontSize: 24,
  },
  closeButton: {
    position: "absolute",
    top: 45,
    right: 20,
    zIndex: 1,
  },
  closeButtonText: {
    fontSize: 30,
  },
  profilePicture: {
    width: '95%',
    aspectRatio: 1,
    backgroundColor: "rgb(200,200,200)",
  },
});

export default MatchesUserCard;