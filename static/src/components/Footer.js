import React, { Component } from 'react';
import Constants from '../Constants';
import Stats from './Stats';
import './Footer.css';


class Footer extends Component {
  render() {
  return (
    <div className="Footer">
      <Stats />

      <div className="Footer-links">
        View source on <a href={Constants.GITHUB_URL}>GitHub</a>
      </div>
    </div>
    );
  }
}

export default Footer;