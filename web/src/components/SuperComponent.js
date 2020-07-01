import React from 'react';

class SuperComponent extends React.Component {
    constructor(props, propValidation = {}) {
      super(props);

      this.propValidation = propValidation;
    }

/*    availablePropTypes = {
        'float' : n => Number(n) === n && n % 1 !== 0,
        'string' : s => typeof s === 'string' 
    }*/

    validatePropTypes() {
        let isValid = false;

        for (const [keyName, validationFunction] of Object.entries(this.propValidation)) {
            if(!validationFunction(this.props[keyName]))
            {
                console.warn(`${this.constructor.name}: ${keyName} did not comply with the validation ${validationFunction}`);
                isValid = false;
            }
            /*if(validationFunction instanceof RegExp && !this.props[keyName].match(validationFunction)) {
                isValid = false;
                console.warn(`${this.constructor.name}: ${keyName} did not match pattern of ${validationFunction}`);
            }
            else if(!(this.props[keyName] instanceof validationFunction))
            {
                isValid = false;
                console.warn(`${this.constructor.name}: ${keyName} is not of the same instance as ${validationFunction}`);
            }
            else if(typeof validationFunction === 'function')
            {
                if(!validationFunction(this.props[keyName]))
                {
                    isValid = false;
                    console.warn(`${this.constructor.name}: ${keyName} is invalid`);
                }
            }*/
        }

        return isValid;
    }
}

SuperComponent.availablePropTypes = {
    'float' : n => Number(n) === n && n % 1 !== 0,
    'string' : s => typeof s === 'string' 
};

export default SuperComponent;