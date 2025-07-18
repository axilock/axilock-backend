import React from "react";
import "./Secretscanrules-edit.css"; // Import the styles
import Button from "react-bootstrap/Button";
import Form from "react-bootstrap/Form";
import Header from "../Common/Header";

const Secretscanedit = () => {
  return (
    <>
      <div className="main">
        <div className="home-container">
          <Header secretscanEdit={true} />
        </div>
      </div>
      <div className="main-banner">
        <div className="banner">
          <div className="banner-title">
            <h4>Secret Scanning Rules / Edit</h4>
          </div>
          <Form>
            <Form.Group
              className="mb-3 name"
              controlId="exampleForm.ControlInput1"
            >
              <Form.Label>Pattern Name</Form.Label>
              <Form.Control type="text" />
            </Form.Group>
            <Form.Group
              className="mb-3 expression"
              controlId="exampleForm.ControlInput1"
            >
              <Form.Label>Secret regular expression *</Form.Label>
              <Form.Control type="text" placeholder="example_[A-Za-z]{40}" />
            </Form.Group>
            <Form.Group
              className="mb-3 matches"
              controlId="exampleForm.ControlTextarea1"
            >
              <Form.Label>Test String- 0 matches</Form.Label>
              <Form.Control
                as="textarea"
                rows={5}
                placeholder="Provide a sample test string to make sure your configuration matches the patterns you except."
              />
            </Form.Group>
            <Button variant="primary grey-btn">Save</Button>
            <Button variant="primary green-btn">Publish</Button>
            <div className="bottom">
              <div className="discard-detail">
                <p>Discard this pattern</p>
              </div>
              <Button type="btn">Discard</Button>
            </div>
          </Form>
        </div>
      </div>
    </>
  );
};

export default Secretscanedit;
