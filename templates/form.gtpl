<html>
    <head>
        <title>OBADA root hash tool</title>
        <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.1/dist/css/bootstrap.min.css">
        <script src="https://unpkg.com/htmx.org@1.4.1" integrity="sha384-1P2DfVFAJH2XsYTakfc6eNhTVtyio74YNA1tc6waJ0bO+IfaexSWXc2x62TgBcXe" crossorigin="anonymous"></script>
    </head>
    <body>
        <div class="container">
            <main>
                <form hx-post="/" hx-target="#obit-hash">

                <div class="row">
                  <div class="form-group col-6">
                    <label for="serial_number_hash">Serial Number Hash:</label>
                    <input type="text" class="form-control" id="serial_number_hash" name="serial_number_hash" required>
                  </div>
                </div>

                <div class="row">
                  <div class="form-group col-6">
                    <label for="manufacturer">Manufacturer:</label>
                    <input type="text" class="form-control" id="manufacturer" name="manufacturer" required>
                  </div>
                </div>

                <div class="row">
                  <div class="form-group col-6">
                      <label for="part_number">Part Number:</label>
                      <input type="text" class="form-control" id="part_number" name="part_number" required>
                  </div>
                </div>

                <div class="row">
                   <div class="form-group col-6">
                       <label for="owner_did">Owner Did:</label>
                       <input type="text" class="form-control" id="owner_did" name="owner_did" required>
                   </div>
                </div>

                <div class="row">
                  <div class="form-group col-6">
                      <label for="obd_did">Obd Did:</label>
                      <input type="text" class="form-control" id="obd_did" name="obd_did" required>
                  </div>
                </div>

                <div class="row">
                  <div class="form-group col-6">
                      <label for="metadata">Metadata:</label>
                      <input type="text" class="form-control" id="metadata" name="metadata" required>
                  </div>
                </div>

                 <div class="row">
                     <div class="form-group col-6">
                          <label for="structured_data">Structured data:</label>
                          <input type="text" class="form-control" id="structured_data" name="structured_data" required>
                     </div>
                 </div>

                  <div class="row">
                      <div class="form-group col-6">
                           <label for="documents">Documents:</label>
                           <input type="text" class="form-control" id="documents" name="documents" required>
                      </div>
                  </div>

                   <div class="row">
                       <div class="form-group col-6">
                            <label for="status">Status:</label>
                            <select type="text" class="form-control" id="status" name="status" required>
                                <option value="FUNCTIONAL">Functional</option>
                                <option value="NON_FUNCTIONAL">Non Functional</option>
                                <option value="DISPOSED">Disposed</option>
                                <option value="STOLEN">Stolen</option>
                                <option value="DISABLED_BY_OWNER">Disabled by owner</option>
                            </select>
                       </div>
                   </div>

                    <div class="row">
                        <div class="form-group col-2">
                          <label for="modified_at">Modified At:</label>
                          <input type="datetime-local" class="form-control" id="modified_at" name="modified_at" min="2021-01-01" max="2030-12-31">
                        </div>
                    </div>

                  <br />

                  <div id="obit-hash"></div>

                  <br />
                  <div class="row">
                    <button type="submit" class="btn btn-primary col-3">Get Root Hash</button>
                  </div>
                </form>
            </main>
        </div>
    </body>
</html>