import React from 'react';

class SuperComponent extends React.Component {
    constructor(props, propValidation = {}) {
      super(props);

      this.propValidation = propValidation;
    }

    validatePropTypes() {
        let isValid = true;

        for (const [keyName, validationEntry] of Object.entries(this.propValidation)) {

            let [validationType, validationFunction] = validationEntry;

            if(!validationFunction(this.props[keyName]))
            {
                console.warn(`${this.constructor.name}: ${keyName} not of type ${validationType}`);
                isValid = false;
            }
        }

        return isValid;
    }
}

SuperComponent.availablePropTypes = {
    'float' : () => ['float', n => Number(n) === n && n % 1 !== 0],
    'string' : () => ['string', s => typeof s === 'string'],
    'choice' : b => [b.split('|'), c => b.includes(c)],
    'range' : (max, min) => [min + ' >= x =< ' + max, d => d <= max && d >= min]
};

export default SuperComponent;