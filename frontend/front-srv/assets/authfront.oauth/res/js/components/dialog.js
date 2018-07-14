import React from 'react';
import Auth from '../containers/Auth';
import OIDCProvider from '../containers/OIDCProvider';
import oidc from '../services/oidc';

class LoginDialog extends React.Component {
    componentDidMount() {
        this.props.signIn()
    }

    render() {
        return <div></div>
    }
}

export default class Dialog extends React.Component {
    render() {
        return (
            <OIDCProvider oidc={oidc}>
                 <Auth>
                 {auth => {
                     if (auth.isAuthed) {
                         this.props.onDismiss()

                         return null
                     }

                     return (
                         <LoginDialog {...auth} />
                     )
                 }}
                 </Auth>
             </OIDCProvider>
         )
    }
}
