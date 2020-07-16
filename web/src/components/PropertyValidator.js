export default function PropertyValidator(props, propValidation = {}) {
    let isValid = true;
    
    for (const [keyName, validationEntry] of Object.entries(propValidation)) {

        let [validationType, validationFunction] = validationEntry;

//        console.log(validationType, validationFunction);

        if(!(keyName in props)){
            console.warn(`missing ${keyName} in configruation, expecting type ${validationType}`);
            isValid = false;
        } else if(!validationFunction(props[keyName])) {
            console.warn(`${keyName} not of type ${validationType}`);
            isValid = false;
        }
    }

    return isValid;
}

export const PropertyValidatorType = {
    'float' : () => ['float', n => Number(n) === n && n % 1 !== 0],
    'string' : () => ['string', s => typeof s === 'string'],
    'choice' : b => [b.split('|'), c => b.includes(c)],
    'range' : (max, min) => [min + ' >= x =< ' + max, d => d <= max && d >= min],
    'regexp' : pattern => ['regexp ' + pattern, input => input.match(pattern)]
};
