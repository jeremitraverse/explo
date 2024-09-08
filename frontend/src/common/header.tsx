import React from 'react'
import './header.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faDesktop } from '@fortawesome/free-solid-svg-icons'
import { faMobileAlt } from '@fortawesome/free-solid-svg-icons'


export default class Header extends React.Component {
    state = { activeViewPort: false }
    constructor() {
        super({})
    }

    setActiveViewPort = () => {
        const newState = !this.state.activeViewPort
        this.setState({"activeViewPort" : newState })
    }

    render() {
        return (
            <div className="header">
                <div className="actions-container">
                    <div id="save">
                        <div className="not-active">SAVE</div>
                    </div>
                    <div id="exit"><div>EXIT</div></div>
                </div>
                <div className="post-name-container">
                    <div className="post-name">Accueil</div>
                    <div className="status">Page - Published</div>
                </div>
                <div className="view-port-toggle-container">
                    <div className="view-port-desktop" onClick={this.setActiveViewPort}>
                        <FontAwesomeIcon icon={faDesktop} />
                    </div>
                    <div className="view-port-mobile">
                        <FontAwesomeIcon icon={faMobileAlt} />
                    </div>
                </div>
            </div>
        )
    }
}