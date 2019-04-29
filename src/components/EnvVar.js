import React from 'react';
import PropTypes from 'prop-types';

const values = {
    marginLeft: '30px'
}

export const EnvVar = ({name}) => {

    if (process.env[name]){
        var displayName = name.replace("REACT_APP_", "");

        return (
            <div>
                <span>{displayName}</span> : 
                <span style={values}>{process.env[name]}</span>
            </div>
        )
    }

    return (<div></div>)
}

EnvVar.propTypes = {
    name: PropTypes.string.isRequired
};

export default EnvVar;