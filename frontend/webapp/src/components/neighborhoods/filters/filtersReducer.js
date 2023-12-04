const filtersReducer = (filters, action) => {
  switch (action.type) {
    case "active": {
      return {
        ...filters,
        [action.name]: { ...filters[action.name], active: action.active },
      };
    }
    case "maxdistance": {
      return {
        ...filters,
        maxdistance: { ...filters["maxdistance"], value: action.value },
      };
    }
    case "agerange": {
      return {
        ...filters,
        agerange: { ...filters["agerange"], [action.name]: action.value },
      };
    }
    default:
      throw Error("Unknown action: " + action.type);
  }
};

export default filtersReducer;
