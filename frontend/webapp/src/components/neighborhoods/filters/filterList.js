const filterList = {
  maxdistance: {
    active: false,
    value: 10,
    name: "Max Distance",
    id: "maxdistance",
    export: function () {
      return [this.value];
    },
  },

  agerange: {
    active: false,
    min: 10,
    max: 50,
    name: "Age Range",
    id: "agerange",
    export: function () {
      return [this.min, this.max];
    },
  },
};

export { filterList };
