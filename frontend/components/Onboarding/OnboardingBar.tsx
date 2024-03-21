import React from 'react';
import { View, StyleSheet } from 'react-native';
import COLORS from '../../colors';

interface OnboardingBarProps {
  selectedCount: number;
}

function OnboardingBar({selectedCount}: OnboardingBarProps) {

  return (
    <View style={[styles.container, {justifyContent: 'center'}]}>
      <View style={{ flexDirection: 'row', width: 346 }}>
        {Array.from({ length: 5 }).map((_, index) => (
          <View
            style={[
              styles.segment,
              index === selectedCount - 1 ? styles.segmentSelected : styles.segmentUnselected,
              { marginRight: index < 4 ? 8 : 0 }, 
            ]}
          />
        ))}
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    width: '100%', 
    height: 21, 
    alignItems: 'center',
  },
  segment: {
    width: 62.8,
    height: 3,
    borderRadius: 100,
  },
  segmentSelected: {
    backgroundColor: COLORS.primary,
  },
  segmentUnselected: {
    backgroundColor: COLORS.disabled,
  },
});

export default OnboardingBar;