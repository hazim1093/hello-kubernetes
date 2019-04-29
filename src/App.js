import React from 'react';
import EnvVar from './components/EnvVar'
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">       
        <h2>
          Hello Kubernetes
        </h2>

        <img src="https://raw.githubusercontent.com/kubernetes/kubernetes/master/logo/logo.png" className="App-logo" alt="logo" />

        <p>
          Environment Variables
        </p>
        <EnvVar name='REACT_APP_HOSTNAME'/>
        <EnvVar name='REACT_APP_VERSION'/>
        <EnvVar name='REACT_APP_KUBERNETES_SERVICE_HOST'/>
        <EnvVar name='REACT_APP_KUBERNETES_NAMESPACE'/>

      </header>
    </div>
  );
}

export default App;
