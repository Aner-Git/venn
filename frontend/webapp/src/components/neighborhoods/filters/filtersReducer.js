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
    default:
      throw Error("Unknown action: " + action.type);
  }
};

export default filtersReducer;
