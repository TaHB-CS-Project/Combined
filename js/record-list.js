
$.getJSON("js/record-list.json", 
    function (data) {
    var info = '';
    var modal_section = '';

    $.each(data, function (key, value) {
        var moreinfo = '';
        var modalcreate_set_name = 'id="c' + value.Record_id + '"';
        var modalcreate_name = '#c' + value.Record_id;
        var modalbody_set_name = 'id="d' + value.Record_id + '"';
        var modalbody_name = '#d' + value.Record_id;
        

        info += '<tr>';

        info += '<td>' + 
            value.Record_id + '</td>';

        info += '<td id=' + value.Hospital_name + '>' + 
            value.Hospital_name + '</td>';

        info += '<td id=' + value.Start_datetime + '>' + 
            value.Start_datetime + '</td>';

/*        info += '<td id=' + value.Medicalemployee_firstname + '>' + 
            value.Medicalemployee_firstname + '</td>';
*/           
        info += '<td id=' + value.Medicalemployee_lastname + '>' + 
            value.Medicalemployee_lastname + '</td>';
        
        info += '<td id=' + value.Procedure_name + '>' + 
            value.Procedure_name + '</td>';
            
 /*       info += '<td id=' + value.Diagnosis_name + '>' + 
            value.Diagnosis_name + '</td>';
 */       
        info += '<td id=' + value.Outcome + '>' + 
            value.Outcome + '</td>';

        info += `<td><button type="button" class="btn btn-primary" data-toggle="modal" data-target="`+ modalcreate_name +`">
                 Moreinfo</button></td>`;

        info += '</tr>'

        modal_section += `<div class="modal fade" ` + modalcreate_set_name + `=" tabindex="-1" role="dialog" aria-labelledby="exampleModalLongTitle" aria-hidden="true">
                <div class="modal-dialog" role="document">
                <div class="modal-content">
                    <div class="modal-header">
                    <h5 class="modal-title" id="exampleModalLongTitle">Record `+ value.Record_id +`</h5>
                    <button type="button" class="close" data-dismiss="modal" aria-label="Close">
                        <span aria-hidden="true">&times;</span>
                    </button>
                    </div>
                    <div class="modal-body text-center" `+ modalbody_set_name +`">

                    </div>
                    <div class="modal-footer">
                    <button type="button" class="btn btn-primary">Edit</button>
                    <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
                    </div>
                    </div>
                </div>
                </div>
                </div>`;
        moreinfo += '<div class="flex-row">'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Record ID: </b></div><div class="col-sm-6">'+ value.Record_id+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Hospital: </b></div><div class="col-sm-6">'+ value.Hospital_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Date: </b></div><div class="col-sm-6">'+ value.Start_datetime+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Dortor Firstname: </b></div><div class="col-sm-6">'+ value.Medicalemployee_firstname+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Doctor Lastname: </b></div><div class="col-sm-6">'+ value.Medicalemployee_lastname+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Procedure: </b></div><div class="col-sm-6">'+ value.Procedure_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Diagnosis: </b></div><div class="col-sm-6">'+ value.Diagnosis_name+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Outcome: </b></div><div class="col-sm-6">'+ value.Outcome+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Special Notes: </b></div><div class="col-sm-6">'+ value.Special_notes+'</div></div>'
        moreinfo += '<br><b>============== Patient Information ==============</b><br><br>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Patient Birthday: </b></div><div class="col-sm-6">'+ value.Patient_birthday+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b>  Patient Sex: </b></div><div class="col-sm-6">'+ value.Patient_sex+'</div></div>'
        moreinfo += '<div class="form-group row"><div class="col-sm-6"><b> Patient Weight: </b></div><div class="col-sm-6">'+ value.Patient_weightlbs+' lbs</div></div>'
        moreinfo += '</div>'




        $('#modal_section').append(modal_section);
        $(modalbody_name).append(moreinfo);

     });
     $('#record_list_datatable').append(info);
    });

    
   
